class Animal19 {
    type: string
    constructor(type: string) {
        this.type = type
    }
  
    @yelling
    greet() {
      console.log(`Hello, I'm a(n) ${this.type}!`)
    }
  }
  
  const typeToYellingMap = {
    cat: 'meow~ meow~'
  }
  
  function yelling(originalMethod: any, context: ClassMethodDecoratorContext) {
    return function(...args: any[]) {
      console.log(typeToYellingMap[this.type])
      originalMethod.call(this, ...args)
    }
  }
  
  const xcat = new Animal19('cat')
  xcat.greet() // meow~ meow~
               // Hello, I'm a(n) cat!