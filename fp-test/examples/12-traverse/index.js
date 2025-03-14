sequence(List.of, Maybe.of(['the facts'])); // [Just('the facts')]

sequence(Task.of, new Map({ a: Task.of(1), b: Task.of(2) })); // Task(Map({ a: 1, b: 2 }))

sequence(IO.of, Either.of(IO.of('buckle my shoe'))); // IO(Right('buckle my shoe'))

sequence(Either.of, [Either.of('wing')]); // Right(['wing'])

sequence(Task.of, left('wing')); // Task(Left('wing'))