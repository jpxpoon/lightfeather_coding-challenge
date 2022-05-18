<template>
  <div id="notification">
    <div class="border rounded">
      <div class="bg-secondary">
        <header class="text-center">
          <h1><strong>Notification Form</strong></h1>
        </header>
      </div>
      <div class="container p-2">
        <form id="notification-form" class="row g-3" @submit.prevent="submitForm">
          <div class="col-6 mb-3">
            <label for="firstname" class="form-label">First Name <span class="text-danger">*</span></label>
            <input type="text" class="form-control" id="firstname" aria-describedby="First name"
              v-model="notificationForm.FirstName">
          </div>
          <div class="col-6 mb-3">
            <label for="lastname" class="form-label">Last Name <span class="text-danger">*</span></label>
            <input type="text" class="form-control" id="lastname" aria-describedby="Last name"
              v-model="notificationForm.LastName">
          </div>
          <div class="col-6 mb-3">
            <input type="checkbox" class="form-check-input" id="email-checkbox"
              v-model="notificationForm.IsNotifyByEmail">
            <label for="email" class="form-label">Email</label>
            <input type="email" class="form-control" id="email" aria-describedby="Email"
              v-model="notificationForm.Email">
          </div>
          <div class="col-6 mb-3">
            <input type="checkbox" class="form-check-input" id="phone-checkbox"
              v-model="notificationForm.IsNotifyByPhone">
            <label for="phonenumber" class="form-label">Phone Number</label>
            <input type="phone" class="form-control" id="phonenumber" aria-describedby="Phone number"
              v-model="notificationForm.Phone">
          </div>
          <div class="col-6 mb-3">
            <label for="supervisors" class="form-label">Supervisors <span class="text-danger">*</span></label>
            <select class="form-control form-select" id="supervisors" v-model="notificationForm.SupervisorID">
              <option value="">Select Supervisors</option>
              <option v-for="item in supervisors" v-bind:value="item.IdentificationNumber">{{ item.Jurisdiction }} - {{
                  item.FirstName
              }},
                {{ item.LastName }}</option>
            </select>
          </div>
          <div class="col-12">
            <button type="submit" class="btn btn-primary">Submit</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import "./notification.css"
import notification from "./notification.js"
export default notification
</script>