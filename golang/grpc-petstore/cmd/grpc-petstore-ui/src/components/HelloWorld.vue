<template>
    <div class="hello">
  <!-- @submit handles any form of submission. -->
  <!-- .prevent keeps the event from bubbling around and doing anything else. -->
  <form @submit.prevent="createPet">
    <label>
      Name:
      <input type="text" v-model="pet.name"/>
    </label>
    <label>
  <select name="petStatus" v-model="pet.status">
    <option value="PetStatusAvailable">Available</option>
    <option value="PetStatusPending">Pending</option>
    <option value="PetStatusSold">Sold</option>
  </select>
    </label>
    <button type="submit">Submit</button>
  </form>
    </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },

  data() {
      return {
		pet: {
			name: "",
			status: "PetStatusUnknown",
		}
	}
  },

  methods : {
            createPet() {
                this.$http.post("http://127.0.0.1:8080/v1/pet", this.pet, { headers: { "content-type": "application/json" } }).then(result => {

                    this.response = result.data;
                });
            }
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
