function auth() {
  return {
    username: "",
    password: "",
    loading: false,
    async authenticate() {
      if (this.loading) {
        return
      }
      if (this.username == "" || this.password == "") {
        toast_error("Username and Password are required")
        return
      }
      this.loading = true
      try {
        const response = await fetch("/api/auth", {
          method: "POST",
          body: JSON.stringify({
            username: this.username,
            password: this.password
          })
        })
        if (response.status == 500) {
          throw new Error()
        }
        if (response.status == 400) {
          toast_error("Invalid Username or Password")
          this.loading = false
          return
        }
        toast_success("Redirecting to dashboard")
        window.location.href = "/"
      } catch (error) {
        this.loading = false
        toast_error("Failed to authenticate. Check logs")
        console.error(error)
      }
    }
  }
}

function app() {
  return {
    get isSelected() {
      return window.location.pathname == ""
    }
  }
}
