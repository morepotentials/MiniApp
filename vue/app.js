import axios from 'axios'

const app = Vue.createApp({
  data(){
    return {
      aSel: false,
      bSel: false,
      cSel: false,
      arr_users: []
    };
  },
  mounted() {
    axios({ method: "GET", "url": "localhost:8080/users" }).then(result => {
        this.arr_users = result;
        console.log(result);
    }, error => {
        console.error(error);
    });
  },
  methods: {
    boxSel(box) {
      if (box === 'A'){
        this.aSel = true;
      } else if (box === 'B'){
        this.bSel = true;
      } else if (box === 'C'){
        this.cSel = true;

      }
    }
  }
});

app.mount("#styling");