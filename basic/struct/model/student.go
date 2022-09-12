package model

/*
模拟工厂模式
*/

//结构体名称首字母大写，跨包可被访问
type Student struct {
	Name string //字段首字母大写，跨包可以被访问
	Age  int
}

//结构体名称首字母小写，跨包不能被访问(但又必须有小写需求)。解决方案-工厂模式
type scores struct {
	Name  string
	Score float64
	level string //首字母小写，则不可以在其它包访问，此时可以提供一个方法
}

/*
因为scores结构体是私有的，只能在本包使用. 通过工厂模式解决!!!
*/
func NewScores(name string, score float64) *scores {
	return &scores{
		Name:  name,
		Score: score,
	}
}

/*
因为scores结构体字段level首字母小写，因此是私有的，只能在本包访问。通过定义方法解决!!!
*/
func (s *scores) GetScoreLevel() string {
	if s.Score > 90 && s.Score < 100 {
		s.level = "good"
	} else {
		s.level = "middle"
	}
	return s.level
}
