package builder

type IBuilder interface {
	setIp(ip string)
	setAccess()
	setPort(por int)
	getRan() Ran
}

func GetBuilder(builderType int) IBuilder {
	if builderType == 0 {
		// 3gpp acesss
		return newRan3gppBuilder()
	} else if builderType == 1 {
		// non 3gpp acess
		return newRanNon3gppBuilder()
	}
	return nil
}
