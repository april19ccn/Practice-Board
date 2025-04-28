// Write your createKitchen function here! ✨
// You'll need to export it so the tests can run it.

// 厨房有多少污垢
type Dirt = number;

// 厨房有多少物品
type Stock = {
    breads: number,
    fruits: number,
    sauces: number,
    vegetables: number,
}

type Budget = number

export type Cleaner = (dirt: Dirt, time?: number) => Dirt

export type Supplier = (expense: number) => Stock

type Recipe = (stock: Stock) => {
    succeeded: false
} | {
    succeeded: true,
    newStock: Stock
}

type Kitchen = {
    announce: () => string,
    clean: (time?: number) => void,
    purchase: (expense: number) => boolean,
    prepare: (recipe: Recipe) => boolean
}

export function createKitchen(budget: Budget, cleaner: Cleaner, supplier: Supplier): Kitchen {
    let dirt: Dirt = 0;
    let stock: Stock = {
        breads: 0,
        fruits: 0,
        sauces: 0,
        vegetables: 0,
    };

    return {
        announce: () => `I have ${dirt} much dirt, ${budget} budget, ${stock.breads} bread(s), ${stock.fruits} fruit(s), ${stock.sauces} sauce(s), and ${stock.vegetables} vegetable(s).`,
        clean: (time?: number) => dirt = cleaner(dirt, time),
        purchase: (expense: number) => { // 题目最开始说满足这个金额 调用supplier， 但后面说这个参数用于清洁，实际test可以推出 这个金额是用于购物
            if (budget < expense) {
                return false
            }

            const ingredients = supplier(expense);

            stock.breads += ingredients.breads;
            stock.fruits += ingredients.fruits;
            stock.sauces += ingredients.sauces;
            stock.vegetables += ingredients.vegetables;

            budget -= expense;

            return true;
        },
        prepare: (recipe: Recipe) => {
            if (dirt >= 100) {
                return false
            }

            dirt++

            let recipeResult = recipe(stock)
            if (recipeResult.succeeded) {
                stock = recipeResult.newStock;
            }

            return recipeResult.succeeded
        }
    }
}