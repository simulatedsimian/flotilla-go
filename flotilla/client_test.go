package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
	"github.com/simulatedsimian/flotilla-go/dock"
)

type RequiredModules struct {
	M1 Matrix
	M2 Matrix
	Touch
	Number
	Dial
}

func TestModule1(t *testing.T) {
	assert := assert.Make(t)

	m := ModuleCommon{}

	m.moduleType = dock.Slider

	ev := Event{}
	ev.ModuleType = dock.Slider
	ev.Params = []int{100}
	ev.EventType = dock.EventConnected
	ev.dockIndex = 1
	ev.Channel = 2

	m.Update(ev)
	assert(m.address).NotNil()
	assert(m.address).Equal(ModuleAddress{dock: ev.dockIndex, channel: ev.Channel})

	ev.EventType = dock.EventDisconnected
	m.Update(ev)
	assert(m.address).IsNil()
}

func TestModule2(t *testing.T) {
	assert := assert.Make(t)

	isCalled := false
	m := ModuleCommon{}

	m.moduleType = dock.Slider
	ev := Event{}
	ev.ModuleType = dock.Slider
	ev.Params = []int{100}
	ev.EventType = dock.EventUpdate
	ev.dockIndex = 1
	ev.Channel = 2

	m.Update(ev)
	assert(isCalled).Equal(false)

	m.OnUpdate(func(params []int) {
		isCalled = true
	})

	m.Update(ev)
	assert(isCalled).Equal(true)

	m.OnUpdate(nil)
	isCalled = false
	m.Update(ev)
	assert(isCalled).Equal(false)
}

func TestAquire(t *testing.T) {
	mustPanic := assert.MustPanic
	assert := assert.Make(t)

	modules := RequiredModules{}

	assert(structMembersToModules(&modules)).Equal(
		[]Module{&Matrix{}, &Matrix{}, &Touch{}, &Number{}, &Dial{}},
	)

	mustPanic(t, func(t *testing.T) {
		structMembersToModules(modules)
	})

	mustPanic(t, func(t *testing.T) {
		structMembersToModules(0)
	})
}

func TestConnectDisconnect(t *testing.T) {
	assert := assert.Make(t)

	e1, e2 := dock.NewPipe().Endpoints()

	client, _ := ConnectToDocksRaw(e1)
	sim := dock.NewSimulator(e2)

	var modules RequiredModules

	client.AquireModules(&modules)
	assert(modules.M1.Connected()).Equal(false)

	sim.Connect(dock.Matrix, 3)
	client.waitForEvent()
	assert(modules.M1.Connected()).Equal(true)

	sim.Disconnect(3)
	client.waitForEvent()
	assert(modules.M1.Connected()).Equal(false)

	e1.Close()

	assert(client.waitForEvent()).HasError()
}
