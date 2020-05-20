import { Controller } from 'stimulus'
import axios from 'axios'

export default class extends Controller {
  loading
  static get targets () {
    return [
      'class', 'name', 'email', 'phone'
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
    const that = this
    try {
      const resp = await axios.post('/api/v1/get-started', req)
      const result = resp.data
      if (result.error) {
        window.alert(result.error)
        that.loading = false
        return
      }
      window.location.href = '/thank-you'
    } catch (error) {
      console.log(error)
      this.loading = false
      window.alert(error)
    }
  }
}
