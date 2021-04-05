Vue.component(
  "user-component",
  {
  props: [
    'user'
  ],
  filters: {
    fullName(user){
      return `${user.first_name} ${user.last_name}`
    }
  },
  template: `
  <div>
    <h4>{{user.id}}</h4>
    <h5>First name is {{user | fullName}}</h5>
  </div>
  `
  } 
);


const app = new Vue({
  el:"#app",
  data: {
      arr_users: [
        {
          id:2, 
          first_name: "Mason", 
          last_name: "child"
        }, 
        {
          id:3, 
          first_name: "Ben", 
          last_name: "Kingsley"
        }
      ],
  },
  mounted(){
    fetch("http://localhost:8080/users")
    .then(response=>response.json())
    .then((data)=>{
      this.arr_users=data
    })
  },
  template: `
  <div>
    <header>
      <h1>Vue Dynamic Styling</h1>
    </header>
    <section id="styling">
      <user-component v-for="user in arr_users" v-bind:user="user"/>

    </section>
  </div>`
})

