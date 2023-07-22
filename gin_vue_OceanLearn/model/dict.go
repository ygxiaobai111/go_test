package model

//字典翻译
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en_us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"` //解释
		Synonym      []string `json:"synonym"`      //同义词
		Antonym      []string `json:"antonym"`      //反义词

		WqxExample []Exapmle `json:"wqx_example"` //例子
		Entry      string    `json:"entry"`
		Type       string    `json:"type"`    //类型
		Related    []any     `json:"related"` //同义词
		Source     string    `json:"source"`  //原字符
	} `json:"dictionary"`
}

//句子
type Exapmle struct {
	ExampleStrings []string `json:"example_strings"`
}
