function handler() {
  return {
    cart_toggled: false,
    cart_count: 10,
    toggleCart() {
      console.log(`cart toggled ${this.cart_toggled}`)
      this.cart_toggled = !this.cart_toggled
      console.log(`cart after toggled ${this.cart_toggled}`)
    }
  }
}
