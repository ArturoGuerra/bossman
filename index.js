#!/usr/bin/env node

const Discord = require("discord.js");
const client = new Discord.Client();

client.on("ready", () => {
  console.log(`Logged in as ${client.user.tag}`);
});

client.on("message", message => {
  if (message.content == '-bossman') {
    vcplay(message, './media1.mp3')
  } else if (message.content == '-musklex') {
    vcplay(message, './media2.mp3')
  } else if (message.content == '-red') {
    vcplay(message, './media3.mp3')
  }
});

async function vcplay(msg, file) {
  if (msg.member.voiceChannel) {
    try {
      vc = await msg.member.voiceChannel.join()
      const dispatcher = vc.playFile(file)
      dispatcher.on("end", end => {
        vc.disconnect()
      })
    } catch (e) {
      console.error(e)
    }
  }
}

client.login(process.env.TOKEN);
