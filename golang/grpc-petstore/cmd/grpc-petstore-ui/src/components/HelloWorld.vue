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
  <select type="number" name="petStatus" v-model.number="pet.status">
    <option value="1">Available</option>
    <option value="2">Pending</option>
    <option value="3">Sold</option>
  </select>
    </label>
    <button type="submit">Submit</button>
  </form>


	<ul v-for="pet in pets">
	<li>
	    {{ pet.id }}: {{ pet.name }}
	</li>
	</ul>

    </div>


</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },

  mounted() {
       this.getPets();
  },

  data() {
      return {
          pets: null,
          pet: {
               name: "",
               status: 0,
            }
	}
  },

  methods : {
	   getPets() {
	               this.$http.get("http://127.0.0.1:8080/v1/pets", null, { headers: { "content-type": "application/json" } }).then(result => {
		   var lines = result.bodyText.split('\n');


			var pets = [];
			// TODO weird hack last line is not good for JSON.parse
			for(var i = 0;i < lines.length-1;i++) {
				var pet;
				pet = JSON.parse(lines[i]);
				pets.push(pet.result);
			}

                    this.pets = pets;
                });
		
	    },

            createPet() {
                this.$http.post("http://127.0.0.1:8080/v1/pets", this.pet, { headers: { "content-type": "application/json" } }).then(result => {

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
