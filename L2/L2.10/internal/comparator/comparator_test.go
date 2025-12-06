package comparator

import "testing"

func TestLexicoComparator(t *testing.T) {
	cmp := &LexicoComparator{}
	if cmp.Compare("a", "b") >= 0 {
		t.Error("a should be less than b")
	}
	if cmp.Compare("b", "a") <= 0 {
		t.Error("b should be greater than a")
	}
	if cmp.Compare("a", "a") != 0 {
		t.Error("a should equal a")
	}
}

func TestLexicoComparatorWithKeyField(t *testing.T) {
	cmp := &LexicoComparator{keyField: 2, delimiter: '\t'}
	line1 := "1\tz"
	line2 := "2\ta"
	if cmp.Compare(line1, line2) <= 0 {
		t.Error("z should be greater than a")
	}
}

func TestNumericComparator(t *testing.T) {
	cmp := &NumericComparator{}
	if cmp.Compare("1", "2") >= 0 {
		t.Error("1 should be less than 2")
	}
	if cmp.Compare("10", "2") <= 0 {
		t.Error("10 should be greater than 2")
	}
	if cmp.Compare("5", "5") != 0 {
		t.Error("5 should equal 5")
	}
}

func TestNumericComparatorWithNonNumeric(t *testing.T) {
	cmp := &NumericComparator{}
	if cmp.Compare("abc", "def") >= 0 {
		t.Error("abc should be less than def lexicographically")
	}
}

func TestNumericComparatorWithKeyField(t *testing.T) {
	cmp := &NumericComparator{keyField: 1, delimiter: '\t'}
	line1 := "10\tother"
	line2 := "2\tother"
	if cmp.Compare(line1, line2) <= 0 {
		t.Error("10 should be greater than 2")
	}
}

func TestMonthComparator(t *testing.T) {
	cmp := &MonthComparator{}
	if cmp.Compare("Jan", "Feb") >= 0 {
		t.Error("Jan should be less than Feb")
	}
	if cmp.Compare("Dec", "Jan") <= 0 {
		t.Error("Dec should be greater than Jan")
	}
	if cmp.Compare("Mar", "Mar") != 0 {
		t.Error("Mar should equal Mar")
	}
}

func TestMonthComparatorWithNonMonth(t *testing.T) {
	cmp := &MonthComparator{}
	if cmp.Compare("abc", "def") >= 0 {
		t.Error("non-month strings should compare lexicographically")
	}
}

func TestHumanComparator(t *testing.T) {
	cmp := &HumanComparator{}
	if cmp.Compare("1K", "2K") >= 0 {
		t.Error("1K should be less than 2K")
	}
	if cmp.Compare("2M", "1M") <= 0 {
		t.Error("2M should be greater than 1M")
	}
	if cmp.Compare("1K", "1024") != 0 {
		t.Error("1K should equal 1024")
	}
}

func TestHumanComparatorWithKeyField(t *testing.T) {
	cmp := &HumanComparator{keyField: 1, delimiter: '\t'}
	line1 := "1M\tother"
	line2 := "512K\tother"
	if cmp.Compare(line1, line2) <= 0 {
		t.Error("1M should be greater than 512K")
	}
}

func TestFactoryCreateLexico(t *testing.T) {
	factory := NewFactory(0, false, false, false, false, '\t')
	cmp := factory.Create()
	if _, ok := cmp.(*LexicoComparator); !ok {
		t.Error("should create LexicoComparator")
	}
}

func TestFactoryCreateNumeric(t *testing.T) {
	factory := NewFactory(0, true, false, false, false, '\t')
	cmp := factory.Create()
	if _, ok := cmp.(*NumericComparator); !ok {
		t.Error("should create NumericComparator")
	}
}

func TestFactoryCreateMonth(t *testing.T) {
	factory := NewFactory(0, false, true, false, false, '\t')
	cmp := factory.Create()
	if _, ok := cmp.(*MonthComparator); !ok {
		t.Error("should create MonthComparator")
	}
}

func TestFactoryCreateHuman(t *testing.T) {
	factory := NewFactory(0, false, false, true, false, '\t')
	cmp := factory.Create()
	if _, ok := cmp.(*HumanComparator); !ok {
		t.Error("should create HumanComparator")
	}
}

func TestIgnoreBlanks(t *testing.T) {
	cmp := &LexicoComparator{ignoreBlanks: true}
	if cmp.Compare("a  ", "a") != 0 {
		t.Error("should ignore trailing blanks")
	}
}
