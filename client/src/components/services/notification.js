import axios from 'axios'

let app = {
    name: 'notification',
    data() {
        return {
            supervisors: [],
            notificationForm: { SupervisorID: "" /* default supervisor*/ }
        }
    },
    methods: {
        resetForm() {
            this.notificationForm = { SupervisorID: "" };
        },
        getSupervisors() {
            axios.get("http://localhost:3001/api/supervisors")
                .then((result) => {
                    this.supervisors = result.data.data;
                    this.supervisors
                        .sort((a, b) => compare(a.FirstName, b.FirstName))
                        .sort((a, b) => compare(a.LastName, b.LastName))
                        .sort((a, b) => compare(a.Jurisdiction, b.Jurisdiction));
                })
                .catch((error) => console.log(error))
        },
        submitForm() {
            axios.post("http://localhost:3001/api/submit", this.notificationForm)
                .then((result) => {
                    if (result.data.message == "") {
                        this.resetForm();
                        console.log(result.data.data);
                        this.$alert("Form Submitted!")
                    } else {
                        this.$alert(result.data.message);
                    }
                })
                .catch((error) => console.log(error))
        }
    },
    mounted() {
        this.getSupervisors();
    }
}

function compare(a, b) {
    if (a < b) { return -1; }
    if (a > b) { return 1; }
    return 0;
}

export default app