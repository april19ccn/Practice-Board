import { OpGetMovieOne } from "@/models/movie";

export const getTitleMovie = async(title: string) => {
    return await OpGetMovieOne({ title });
}