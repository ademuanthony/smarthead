import { Controller } from 'stimulus'
import axios from 'axios'

export default class extends Controller {
  loading
  static get targets () {
    return [
      'class', 'name', 'email', 'phone', 'message'
    ]
  }

  async getStarted (e) {
    e.preventDefault()
    if (this.classTarget.value === '' || this.nameTarget.value === '' ||
      this.phoneTarget.value === '' || this.emailTarget.value === '') {
      window.alert('Please select your class and enter your name, email and phone number to continue')
      return
    }

    if (this.loading) return
    this.loading = true

    const req = {
      'class_id': this.classTarget.value,
      'name': this.nameTarget.value,
      'phone': this.phoneTarget.value,
      'email': this.emailTarget.value
    }
    try {
      await axios.post('/api/v1/get-started', req)
      window.location.href = '/thank-you'
    } catch (error) {
      this.loading = false
      if (error.response) {
        const data = error.response.data
        this.messageTarget.textContent = data.details
        window.alert(data.details)
        return
      }
      window.alert('An unknown error has caused the request to fail')
    }
  }
}
