package mypattern

const split  = "--"

type ComputerAbsFactory interface {

	Produce(val interface{}) Computer
}

type Computer interface {
	Configuration() string
	SetCpu(string)
	SetOs(string)
	SetMemory(string)
	SetHardDisk(string)
}

type GetProperty interface {
	getCpu() string
	getOs() string
	getMemory() string
	getHardDisk() string
}

type Mac struct {

	cpu string
	os string
	memory string
	hardDisk string
}

type ThinkPad struct {

	cpu string
	os string
	memory string
	hardDisk string
}

func (thinkpad *ThinkPad) SetCpu(val string){
	if thinkpad!=nil {
		thinkpad.cpu = val
	}
}

func (thinkpad *ThinkPad) SetOs(val string){
	if thinkpad!=nil {
		thinkpad.os = val
	}
}

func (thinkpad *ThinkPad) SetMemory(val string){
	if thinkpad!=nil {
		thinkpad.memory = val
	}
}

func (thinkpad *ThinkPad) SetHardDisk(val string){
	if thinkpad!=nil {
		thinkpad.hardDisk = val
	}
}


func (mac *Mac) SetCpu(val string){
	if mac!=nil {
		mac.cpu = val
	}
}

func (mac *Mac) SetOs(val string){
	if mac!=nil {
		mac.os = val
	}
}

func (mac *Mac) SetMemory(val string){
	if mac!=nil {
		mac.memory = val
	}
}

func (mac *Mac) SetHardDisk(val string){
	if mac!=nil {
		mac.hardDisk = val
	}
}

func (mac Mac) Configuration() string {

	return "Hi, I'm Mac"+split+mac.cpu+split+mac.os+split+mac.memory+split+mac.hardDisk;
}

func (thinkpad ThinkPad) Configuration() string {

	return "Hi, I'm ThinkPad"+split+thinkpad.cpu+split+thinkpad.os+split+thinkpad.memory+split+thinkpad.hardDisk;
}

func Produce(val interface{}) Computer {

	switch val.(type) { //多选语句switch
	case Mac:
		mac := new(Mac)
		mac.cpu = "Intel Core i7"
		mac.os = "macOS Sierra"
		mac.hardDisk = "500g"
		mac.memory = "16g"

		return mac
	case ThinkPad:
		thinkPad :=new(ThinkPad)
		thinkPad.cpu = "Intel Core i5"
		thinkPad.os = "win7"
		thinkPad.hardDisk = "500g"
		thinkPad.memory = "16g"

		return thinkPad
	default:
		return nil
	}

}




