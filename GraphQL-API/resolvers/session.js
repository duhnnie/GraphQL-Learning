module.exports = {
  speakers: async (session, args, { dataSources }, info) => {
    const allSpeakers = await dataSources.speakerAPI.getSpeakers() || [];
    const sessionSpeakersIds = session.speakers.map(({ id }) => id);

    return allSpeakers.filter((speaker) => sessionSpeakersIds.includes(speaker.id));
  },
};
