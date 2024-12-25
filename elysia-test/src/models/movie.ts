import visitDB from "../config/connectDB";

export const OpGetMovieOne = async (condition: any) => {
    // try {
    //     await client.connect()
    //     const database = client.db('sample_mflix');
    //     const movies = database.collection('movies');

    //     // Query for a movie that has the title 'Back to the Future'
    //     // const query = { title: 'The Favourite' };
    //     const query = condition
    //     const movie = await movies.findOne(query);

    //     console.log(movie);
    //     return movie
    // } finally {
    //     // Ensures that the client will close when you finish/error
    //     await client.close();
    // }

    return await visitDB(async (client: any)=> {
        const database = client.db('sample_mflix');
        const movies = database.collection('movies');

        // Query for a movie that has the title 'Back to the Future'
        // const query = { title: 'The Favourite' };
        const query = condition
        const movie = await movies.findOne(query);

        console.log(movie);
        return movie
    })
};