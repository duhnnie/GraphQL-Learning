const { DataSource } = require('apollo-datasource');
const _ = require('lodash');
const sessions = require('../data/sessions.json');

class SessionAPI extends DataSource {
  constructor() {
    super();
  }

  initialize(config) {}

  getSessions(args) {
    return _.filter(sessions, args);
  }

  getSessionById(id) {
    return sessions.find((session) => session.id === parseInt(id));
  }

  toggleFavorite(id) {
    const session = sessions.find((session) => session.id === parseInt(id));

    if (session) {
      session.favorite = !session.favorite;
    }

    return session;
  }

  addSession(session) {
    session.id = 12;
    sessions.push(session);

    return session;
  }
}

module.exports = SessionAPI;
