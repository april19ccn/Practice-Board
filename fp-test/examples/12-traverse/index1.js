import { IO, Maybe, Either, Identity, left, compose, chain, map, sequence, traverse } from "../../utils/support.js";

// 同一律 (Identity)
const identity1 = compose(sequence(Identity.of), map(Identity.of));
const identity2 = Identity.of;

// test it out with Right
console.log(identity1(Either.of('stuff')));
// Identity(Right('stuff'))

console.log(identity2(Either.of('stuff')));
// Identity(Right('stuff'))


// Right { '$value': Identity { '$value': 'stuff' } }
// Right(Identity('stuff')).sequence(Identity.of)
// => Right(Identity('stuff')).traverse(of, identity)
// => identity(Identity('stuff')).map(Either.of)
// => Identity(Right('stuff'))


// 组合 (Composition)
class Compose { // CH.8
    constructor(fgx) {
        this.getCompose = fgx;
    }

    static of(fgx) {
        return new Compose(fgx);
    }

    map(fn) {
        return new Compose(map(map(fn), this.getCompose));
    }
}

const comp1 = compose(sequence(Compose.of), map(Compose.of));
const comp2 = (Fof, Gof) => compose(Compose.of, map(sequence(Gof)), sequence(Fof));

// Test it out with some types we have lying around
console.log(comp1(Identity.of(Either.of([true]))));
// Compose(Right([Identity(true)]))

console.log(comp2(Either.of, Array)(Identity.of(Either.of([true]))));
// Compose(Right([Identity(true)]))

// comp1(Identity.of(Either.of([true]))) 
// = compose(sequence(Compose.of), map(Compose.of))(Identity.of(Either.of([true]))) 
// ----> Identity.of(Either.of([true])).map(Compose.of) = Identity.of(Compose.of(Either.of([true])))
// ----> Identity(Compose(Right([true]))).sequence(Compose.of) 
        // = Identity(Compose(Right([true]))).traverse(Compose.of, identity) 
        // = identity(Compose(Right([true]))).map(Identity.of)
        // = new Compose( map ( map(Identity.of), Right([true]) ) )
        // = Compose(Right([Identity(true)]))

// comp2(Either.of, Array)(Identity.of(Either.of([true])))
// = (compose(Compose.of, map(sequence(Array)), sequence(Either.of)))(Identity.of(Either.of([true])))
// ----> Identity.of(Either.of([true])).sequence(Either.of) 
        // = Identity.of(Either.of([true])).traverse(Either.of, identity) 
        // = identity(Either.of([true])).map(Identity.of)
        // = Right(Identity([true]))
// ----> Right(Identity[true]).map(sequence(Array))
        // = Right(Identity([true]).sequence(Array))
        // = Right(Identity([true]).traverse(Array, identity))
        // = Right(identity([true]).map(Identity.of))
        // = Right([Identity(true)])
// ----> Compose.of(Right([Identity(true)]))


// 自然 (Naturality)
const natLaw1 = (of, nt) => compose(nt, sequence(of));
const natLaw2 = (of, nt) => compose(sequence(of), map(nt));

// test with a random natural transformation and our friendly Identity/Right functors.

// maybeToEither :: Maybe a -> Either () a
const maybeToEither = (x) => (x.$value ? Either.of(x.$value) : left());

console.log(natLaw1(Maybe.of, maybeToEither)(Identity.of(Maybe.of('barlow one'))));
// Right(Identity('barlow one'))

console.log(natLaw2(Either.of, maybeToEither)(Identity.of(Maybe.of('barlow one'))));
// Right(Identity('barlow one'))