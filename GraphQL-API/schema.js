const { gql } = require('apollo-server');

module.exports = gql`
type Query {
  sessions(
    id: ID
    title: String
    description: String
    startsAt: String
    endsAt: String
    room: String
    day: String
    format: String
    track: String
    level: String
  ): [Session]
  sessionById(id: Int): Session
  speakers: [Speaker]
  speakerById(id: String): Speaker
}

type Mutation {
  toggleSessionFavorite(id: ID): Session
  addSession(session: SessionInput): Session
}

input SessionInput {
  title: String!
  description: String
  startsAt: String
  endsAt: String
  room: String
  day: String
  format: String
  track: String
  level: String
  favorite: Boolean
}

type Speaker {
  id: ID!
  bio: String
  name: String
  sessions: [Session]
}

type Session {
  id: ID!
  title: String!
  description: String
  startsAt: String
  endsAt: String
  room: String
  day: String
  format: String
  track: String @deprecated(reason: "A tag based system will be used in the future.")
  level: String
  favorite: Boolean
  speakers: [Speaker]
}
`;