<template>
    <v-form v-model="valid" ref="form" lazy-validation>
        <v-text-field
                label="E-mail"
                type="email"
                v-model="email"
                :rules="emailRules"
                required
        ></v-text-field>

        <v-text-field
                label="Password"
                type="password"
                v-model="password"
                :rules="passwordRules"
                required
        ></v-text-field>

        <v-btn
                @click="submit"
                :disabled="!valid"
        >
            login
        </v-btn>
    </v-form>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      email: '',
      emailRules: [
        (v) => !!v || 'E-mail is required',
        (v) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
      ],
      password: '',
      passwordRules: [
        (v) => !!v || 'Password is required'
      ]
    }),
    methods: {
      submit () {
        if (this.$refs.form.validate()) {
          // Native form submission is not yet supported
//          axios.post('/api/submit', {
//            name: this.name,
//            email: this.email,
//            select: this.select,
//            checkbox: this.checkbox
//          });
        }
      },
      clear () {
        this.$refs.form.reset();
      }
    },
    title () {
      return 'Login';
    }
  };
</script>