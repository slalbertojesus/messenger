<template>
  <v-container>
    <section>
      <h1>gRPC Client</h1>
      <input
        v-model="inputField"
        v-on:keyup.enter="YesHello()"
        placeholder="Mandar mensaje"
      />
      <v-btn @click="YesHello()">Mandar</v-btn>
    </section>
  </v-container>
</template>

<script>
import { PingMessage } from "@/grpc-client/api_pb.js";
import { PingClient } from "@/grpc-client/api_grpc_web_pb.js";

export default {
  name: "HelloWorld",
  created: function() {
    this.client = new PingClient("http://localhost:7777", null, null);
  },
  methods: {
    YesHello() {
      let request = new PingMessage();
      request.setGreeting(this.inputField);
      this.client.yesHello(request, {}, err => {
        if (err) {
          console.log(err);
        }
        this.inputField = "";
      });
    }
  },
  data: () => ({
    inputField: ""
  })
};
</script>
