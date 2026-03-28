import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";
import db from "./_db.js";

// types
import { typeDefs } from "./schema.js";

const resolvers = {
    Query: {
        reviews() {
            return db.reviews
        },
        games() {
            return db.games
        },
        authors() {
            return db.authors
        },
        review(_, args, ctx) {
            return db.reviews.find(review => review.id === args.id)
        },
        game(_, args, ctx) {
            return db.games.find(game => game.id === args.id)
        },
        author(_, args, ctx) {
            return db.authors.find(author => author.id === args.id)
        }
    },
    Game: {
        reviews(parent) {
            return db.reviews.filter(i => i.game_id === parent.id)
        }
    },
    Author: {
        reviews(parent) {
            return db.reviews.filter(i => i.game_id === parent.id)
        }
    },
    Review: {
        game(parent) {
            return db.games.find(g => g.id === parent.game_id)
        },
        author(parent) {
            return db.authors.find(a => a.id === parent.author_id)
        }
    },
    Mutation: {
        addGame(parent, args) {
            const game = {
                ...args.input,
                id: Math.floor(Math.random() * 10000)
            }

            db.games.push(game)
            return game
        },
        updateGame(parent, args) {
            db.games = db.games.map(g => {
                if(g.id === args.id) {
                    return {
                        ...g,
                        ...args.edits
                    }
                }

                return g
            })

            return db.games.find(g => g.id === args.id)
        },
        deleteGame(parent, args) {
            db.games = db.games.filter(i => i.id !== args.id)
            return db.games
        }
    }
}

// server setup
const server = new ApolloServer({
    typeDefs,
    resolvers,
})

const { url } = await startStandaloneServer(server, {
    listen: { port: 4000 }
})

console.log(`Server ready at port ${4000}`)