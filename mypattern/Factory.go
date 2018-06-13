package pattern

const split  = "--"

type ComputerAbsFactory interface {

	Produce(val interface{}) Computer
}

type Computer interface {
	Configuration() string
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

func (mac Mac) Configuration() string {
	
	return "Hi, I'm Mac"+split+mac.cpu+split+mac.os+split+mac.memory+split+mac.hardDisk;
}

func (thinkpad ThinkPad) Configuration() string {

	return "Hi, I'm Mac"+split+thinkpad.cpu+split+thinkpad.os+split+thinkpad.memory+split+thinkpad.hardDisk;
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




