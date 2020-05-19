import { Controller } from 'stimulus'
import axios from 'axios'

export default class extends Controller {
  static get targets () {
    return [
      'username', 'password'
    ]
  }

  async subscribe () {
    if (this.list.length === 0) {
      window.alert('Please add a subject, period amd class to continue')
      return
    }

    if (this.loading) return
    this.loading = true

    const req = {
      'count': this.list.length,
      'subject_id': this.list[0].subject_id,
      'class_id': this.list[0].class_id
    }
    const that = this
    const resp = await axios.post('/payments/initiate', req)
    const result = resp.data
    if (result.error) {
      window.alert(result.error)
      that.loading = false
      return
    }
    // eslint-disable-next-line no-undef
    var handler = PaystackPop.setup({
      key: 'pk_live_a669ccbde1c4d5509b6565af1131f87ea1af5ab6',
      // key: 'pk_test_6e0424ff8e08138c2fcce3f4f2b05052c4b3f77c',
      email: result.student.parent_email,
      amount: result.amount,
      currency: 'NGN',
      ref: result.id,
      metadata: {
        custom_fields: [
          {
            display_name: 'Mobile Number',
            variable_name: 'mobile_number',
            value: result.student.parent_phone
          }
        ]
      },
      callback: async function (response) {
        let items = []
        for (let i = 0; i < that.list.length; i++) {
          items.push({
            'subject_id': that.list[i].subject_id,
            'class_id': that.list[i].class_id
          })
        }
        const resp = await axios.post(`/payments/${response.reference}/update-status`, { 'items': items })
        const result = resp.data
        if (result.error) {
          that.loading = false
          window.alert(result.error)
        }
        window.alert('Success. Subscription successful, check the subscription table for your lesson\' start date')
        // todo: mark payment as succeded and reload the page with a pop for selecting subject and period
        window.location.reload()
      },
      onClose: function () {
        that.loading = false
        window.alert('window closed')
      }
    })
    handler.openIframe()
  }
}
