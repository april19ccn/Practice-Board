package q

import p "example/learn/ch4/04-struct/03-member-p"

// cannot refer to unexported field a in struct literal of type p.T
// cannot refer to unexported field b in struct literal of type p.T 不能在 p.T 类型的 struct literal 中引用未导出字段 b
// var _ = p.T{a: 1, b: 2} // compile error: can't reference a, b

// implicit assignment to unexported field a in struct literal of type p.T 隐式分配给 p.T 类型的 struct literal 的未导出字段 a
// implicit assignment to unexported field b in struct literal of type p.T
// var _ = p.T{1, 2}       // compile error: can't reference a, b

// implicit assignment to unexported field y in struct literal of type p.Z
// var _ = p.Z{1, 2}

// too few values in struct literal of type p.Z
// var _ = p.Z{1}

var _ = p.Z{X: 1}
