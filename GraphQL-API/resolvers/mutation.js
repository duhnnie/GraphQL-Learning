module.exports = {
  toggleSessionFavorite: (parent, { id }, { dataSources }, info) => {
    return dataSources.sessionAPI.toggleFavorite(id);
  },
  addSession: (parent, { session }, { dataSources }, info) => {
    return dataSources.sessionAPI.addSession(session);
  },
};
