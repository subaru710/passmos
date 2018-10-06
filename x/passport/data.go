package passport

type PersonalData struct {
	Name   string `json:name`
	Age    int8   `json:age`
	Gender int8   `json:gender`
}

func (pd PersonalData) GetName() string {
	return pd.Name
}

func (pd *PersonalData) SetName(name string) {
	pd.Name = name
}

func (pd PersonalData) GetAge() int8 {
	return pd.Age
}

func (pd *PersonalData) SetAge(age int8) {
	pd.Age = age
}

func (pd PersonalData) GetGender() int8 {
	return pd.Gender
}

func (pd *PersonalData) SetGender(gender int8) {
	pd.Gender = gender
}
