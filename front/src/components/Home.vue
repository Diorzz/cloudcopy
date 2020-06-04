<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <el-row type="flex" class="row-bg" justify="center">
      <el-col :xs="24" :sm="16" :md="16">
        <el-input v-model="authCode" placeholder="请输入授权码"></el-input>
      </el-col>
    </el-row>

    <div v-on:click="say" class="b_icon">
      <i class="el-icon-right icon-with-border"></i>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      msg: 'Easy your copy between multi devices',
      authCode: ''
    }
  },

  methods: {
    say(message) {
      this.axios.get(
        'http://127.0.0.1:8888',
        ''
      ).then(res => {
        console.log(res)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.b_icon {
  margin-top: 20px;
}
.icon-with-border {
  border-radius: 30px;
  border: 1px solid;
  font-size: 40px;
}
.icon-with-border:hover {
  color: #42b983;
  cursor: pointer;
}
</style>
