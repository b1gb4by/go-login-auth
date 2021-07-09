<template>
  <main class="form-register">
    <form @submit.prevent="submit">
      <h1 class="h3 mb-3 fw-normal">Please Register</h1>
      <input
        v-model="firstName"
        class="form-control"
        placeholder="First Name"
        required>
      <input
        v-model="lastName"
        class="form-control"
        placeholder="Last Name"
        required>
 
      <input
        v-model="email"
        type="email"
        class="form-control"
        placeholder="Email"
        required>
 
      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="Password"
        required>
      <input
        v-model="passwordConfirm"
        type="password"
        class="form-control"
        placeholder="Password Confirm"
        required>
 
      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit">Submit</button>
    </form>
  </main>
</template>
 
<script>
export default {
    name: "Register"
}
</script>
 
<style>
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}
 
.form-register .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-register .form-control:focus {
  z-index: 2;
}
.form-register input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-register input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
 