package app

import (
	"context"
	"net/url"
	"sort"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

type Engine struct {
	// The number of frame per seconds.
	FrameRate int

	// The page.
	Page Page

	// Reports whether the engine runs on server-side.
	IsServerSide bool

	// The storage use as local storage.
	LocalStorage BrowserStorage

	// The storage used as session storage.
	SessionStorage BrowserStorage

	// The function used to resolve static resource paths.
	StaticResourceResolver func(string) string

	// The body of the page.
	Body HTMLBody

	// The action handlers that are not associated with a component and are
	// executed asynchronously.
	ActionHandlers map[string]ActionHandler

	initOnce             sync.Once
	startOnce            sync.Once
	closeOnce            sync.Once
	wait                 sync.WaitGroup
	componentUpdateMutex sync.RWMutex

	dispatches           chan Dispatch
	componentUpdates     map[Composer]bool
	componentUpdateQueue []componentUpdate
	deferables           []Dispatch
	actions              actionManager
	states               *store
	isFirstMount         bool
}

func (e *Engine) Context() Context {
	return makeContext(e.Body)
}

func (e *Engine) Dispatch(d Dispatch) {
	if d.Source == nil {
		d.Source = e.Body
	}
	e.dispatches <- d
}

func (e *Engine) Emit(src UI, fn func()) {
	e.Dispatch(Dispatch{
		Mode:   Next,
		Source: src,
		Function: func(ctx Context) {
			if fn != nil {
				fn()
			}

			e.componentUpdateMutex.RLock()
			compo := getComponent(src)
			if canUpdate, ok := e.componentUpdates[compo]; ok && !canUpdate {
				e.componentUpdateMutex.RUnlock()
				return
			}
			e.componentUpdateMutex.RUnlock()

			for c := compo; c != nil; c = getComponent(c.getParent()) {
				e.addComponentUpdate(c)
			}
		},
	})
}

func (e *Engine) Handle(actionName string, src UI, h ActionHandler) {
	e.actions.handle(actionName, false, src, h)
}

func (e *Engine) Post(a Action) {
	e.Async(func() {
		e.actions.post(a)
	})
}

func (e *Engine) SetState(state string, v any, opts ...StateOption) {
	e.states.Set(state, v, opts...)
}

func (e *Engine) GetState(state string, recv any) {
	e.states.Get(state, recv)
}

func (e *Engine) DelState(state string) {
	e.states.Del(state)
}

func (e *Engine) ObserveState(state string, elem UI) Observer {
	return e.states.Observe(state, elem)
}

func (e *Engine) Async(fn func()) {
	e.wait.Add(1)
	go func() {
		fn()
		e.wait.Done()
	}()
}

func (e *Engine) Wait() {
	e.wait.Wait()
}

func (e *Engine) Consume() {
	for {
		e.Wait()

		select {
		case d := <-e.dispatches:
			e.handleDispatch(d)

		default:
			e.handleFrame()
			return
		}
	}
}

func (e *Engine) ConsumeNext() {
	e.Wait()
	e.handleDispatch(<-e.dispatches)
	e.handleFrame()
}

func (e *Engine) Close() {
	e.closeOnce.Do(func() {
		e.Consume()
		e.Wait()

		dismount(e.Body)
		e.Body = nil
		e.states.Close()
	})
}

func (e *Engine) PreRender() {
	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			ctx.Src().preRender(e.Page)
		},
	})
}

func (e *Engine) Mount(v UI) {
	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			if e.isFirstMount {
				if err := e.Body.(*htmlBody).replaceChildAt(0, v); err != nil {
					panic(errors.New("mounting first ui element failed").Wrap(err))
				}

				e.isFirstMount = false
				return
			}

			if firstChild := e.Body.getChildren()[0]; canUpdate(firstChild, v) {
				if err := update(firstChild, v); err != nil {
					panic(errors.New("mounting ui element failed").Wrap(err))
				}
				return
			}

			if err := e.Body.(*htmlBody).replaceChildAt(0, v); err != nil {
				panic(errors.New("mounting ui element failed").Wrap(err))
			}
		},
	})
}

func (e *Engine) Nav(u *url.URL) {
	if p, ok := e.Page.(*requestPage); ok {
		p.ReplaceURL(u)
	}

	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			ctx.Src().onComponentEvent(nav{})
		},
	})
}

func (e *Engine) AppUpdate() {
	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			ctx.Src().onComponentEvent(appUpdate{})
		},
	})
}

func (e *Engine) AppInstallChange() {
	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			ctx.Src().onComponentEvent(appInstallChange{})
		},
	})
}

func (e *Engine) AppResize() {
	e.Dispatch(Dispatch{
		Mode:   Update,
		Source: e.Body,
		Function: func(ctx Context) {
			ctx.Src().onComponentEvent(resize{})
		},
	})
}

func (e *Engine) init() {
	e.initOnce.Do(func() {
		if e.FrameRate <= 0 {
			e.FrameRate = 60
		}

		if e.Page == nil {
			u, _ := url.Parse("https://test.go-app.dev")
			e.Page = &requestPage{url: u}
		}

		if e.LocalStorage == nil {
			e.LocalStorage = NewMemoryStorage()
		}

		if e.SessionStorage == nil {
			e.SessionStorage = NewMemoryStorage()
		}

		if e.StaticResourceResolver == nil {
			e.StaticResourceResolver = func(path string) string {
				return path
			}
		}

		if e.Body == nil {
			body := Body().privateBody(Div())
			if err := mount(e, body); err != nil {
				panic(errors.New("mounting engine default body failed").Wrap(err))
			}
			e.Body = body
		}

		e.dispatches = make(chan Dispatch, 4096)
		e.componentUpdates = make(map[Composer]bool)
		e.componentUpdateQueue = make([]componentUpdate, 0, 32)
		e.deferables = make([]Dispatch, 32)
		e.states = newStore(e)
		e.isFirstMount = true

		for actionName, handler := range e.ActionHandlers {
			e.actions.handle(actionName, true, e.Body, handler)
		}
	})
}

func (e *Engine) getCurrentPage() Page {
	return e.Page
}

func (e *Engine) getLocalStorage() BrowserStorage {
	return e.LocalStorage
}

func (e *Engine) getSessionStorage() BrowserStorage {
	return e.SessionStorage
}

func (e *Engine) isServerSide() bool {
	return e.IsServerSide
}

func (e *Engine) resolveStaticResource(path string) string {
	return e.StaticResourceResolver(path)
}

func (e *Engine) addComponentUpdate(c Composer) {
	if c == nil || !c.Mounted() {
		return
	}

	e.componentUpdates[c] = true
}

func (e *Engine) removeComponentUpdate(c Composer) {
	delete(e.componentUpdates, c)
}

func (e *Engine) preventComponentUpdate(c Composer) {
	e.componentUpdateMutex.Lock()
	defer e.componentUpdateMutex.Unlock()

	e.componentUpdates[c] = false
}

func (e *Engine) addDeferable(d Dispatch) {
	e.deferables = append(e.deferables, d)
}

func (e *Engine) start(ctx context.Context) {
	e.startOnce.Do(func() {
		frameDuration := time.Second / time.Duration(e.FrameRate)
		currentFrameDuration := frameDuration
		frames := time.NewTicker(frameDuration)

		cleanups := time.NewTicker(time.Minute)
		defer cleanups.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case d := <-e.dispatches:
				if currentFrameDuration != frameDuration {
					currentFrameDuration = frameDuration
					frames.Reset(currentFrameDuration)
				}
				e.handleDispatch(d)

			case <-frames.C:
				e.handleFrame()
				if len(e.dispatches) == 0 {
					if currentFrameDuration < time.Hour {
						currentFrameDuration *= 2
					}
					frames.Reset(currentFrameDuration)
				}

			case <-cleanups.C:
				e.actions.closeUnusedHandlers()
				e.states.Cleanup()
			}
		}
	})
}

func (e *Engine) handleDispatch(d Dispatch) {
	switch d.Mode {
	case Update:
		d.do()
		e.addComponentUpdate(getComponent(d.Source))

	case Defer:
		e.deferables = append(e.deferables, d)

	case Next:
		d.do()
	}
}

func (e *Engine) handleFrame() {
	e.handleComponentUpdates()
	e.handleDeferables()
}

func (e *Engine) handleComponentUpdates() {
	e.componentUpdateMutex.Lock()
	defer e.componentUpdateMutex.Unlock()

	for c, canUpdate := range e.componentUpdates {
		if c.Mounted() && canUpdate {
			e.componentUpdateQueue = append(e.componentUpdateQueue, componentUpdate{
				component: c,
				priority:  getComponentPriority(c),
			})
		}
	}

	sort.Slice(e.componentUpdateQueue, func(i, j int) bool {
		return e.componentUpdateQueue[i].priority < e.componentUpdateQueue[j].priority
	})

	for i, u := range e.componentUpdateQueue {
		if _, ok := e.componentUpdates[u.component]; !ok || !u.component.Mounted() {
			e.removeComponentUpdate(u.component)
			e.componentUpdateQueue[i] = componentUpdate{}
			continue
		}

		if err := u.component.updateRoot(); err != nil {
			panic(err)
		}
		e.removeComponentUpdate(u.component)
		e.componentUpdateQueue[i] = componentUpdate{}
	}

	e.componentUpdateQueue = e.componentUpdateQueue[:0]
}

func (e *Engine) handleDeferables() {
	for i := range e.deferables {
		e.deferables[i].do()
		e.deferables[i] = Dispatch{}
	}
	e.deferables = e.deferables[:0]
}

func getComponent(n UI) Composer {
	for node := n; node != nil; node = node.getParent() {
		if c, isCompo := node.(Composer); isCompo {
			return c
		}
	}
	return nil
}

func getComponentPriority(c Composer) int {
	depth := 1
	for parent := c.getParent(); parent != nil; parent = parent.getParent() {
		depth++
	}
	return depth
}

type componentUpdate struct {
	component Composer
	priority  int
}
