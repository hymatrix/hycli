package tnode

import (
	"github.com/hymatrix/hymx/common"
	vmmSchema "github.com/hymatrix/hymx/vmm/schema"
)

var log = common.NewLog("tnode")

func Spawn(env vmmSchema.Env) (vm vmmSchema.Vm, err error) {
	vmd, err := New(env)
	if err != nil {
		return
	}

	log.Info("spawn process success", "pid", env.Meta.Pid, "from", env.Meta.AccId)
	return vmd, nil
}

type Tnode struct {
	pid string
	Env vmmSchema.Env
	
	// Insert your code here
}

func New(env vmmSchema.Env) (*Tnode, error) {
	return &Tnode{
		pid: env.Meta.Pid,
		Env: env,
	}, nil
}

func (v *Tnode) Apply(from string, meta vmmSchema.Meta) (res vmmSchema.Result) {
	return
}

func (v *Tnode) Checkpoint() (data string, err error) {
	return
}

func (v *Tnode) Restore(data string) error {
	return nil
}

func (v *Tnode) Close() error {
	return nil
}
