<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <p>

        Go to system control, please <a href="/auth/login" target="_blank" >sign in</a>.
    </p>
    <h3>What is included</h3>

    <ul>
      <li><a href="https://github.com/lambda-platform/lambda" target="_blank" rel="noopener">Lambda</a></li>
        <li><a href="hhttps://github.com/lambda-platform/puzzle" target="_blank" rel="noopener">Puzzle</a></li>
      <li><a href="https://github.com/lambda-platform/agent" target="_blank" rel="noopener">Agent</a></li>
      <li><a href="https://github.com/lambda-platform/datasource" target="_blank" rel="noopener">Data Source</a></li>
      <li><a href="https://github.com/lambda-platform/krud" target="_blank" rel="noopener">Krud</a></li>
      <li><a href="https://github.com/lambda-platform/dataform" target="_blank" rel="noopener">Data From</a></li>
      <li><a href="https://github.com/lambda-platform/datagrid" target="_blank" rel="noopener">Data Grid</a></li>
      <li>App example</li>
      <li>App Admin  example (optional)</li>
    </ul>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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
</style>
