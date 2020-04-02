import { Controller } from 'stimulus'
import axios from 'axios'

export default class extends Controller {
  static get targets () {
    return [
      'class', 'subject', 'period'
    ]
  }

  async initiate () {
    const subject = this.subjectTarget.value
    const period = this.periodTarget.value
    const classID = this.classTarget.value
    if (subject === '' || period === '' || classID === '') {
      window.alert('Subject, period amd class is required')
      return
    }

    if (this.loading) return
    this.loading = true

    const req = {
      'subject_id': subject,
      'period_id': period,
      'class_id': classID
    }
    const resp = await axios.post('/payments/initiate', req)
    const result = resp.data
    if (result.error) {
      window.alert(result.error)
      return
    }
    // eslint-disable-next-line no-undef
    var handler = PaystackPop.setup({
      key: 'pk_live_a669ccbde1c4d5509b6565af1131f87ea1af5ab6',
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
        const resp = await axios.get(`/payments/${response.reference}/update-status`)
        const result = resp.data
        if (result.error) {
          window.alert(result.error)
        }
        window.alert('Success. Subscription successful, check the subscription table for your lesson start date')
        // todo: mark payment as succeded and reload the page with a pop for selecting subject and period
        window.location.reload()
      },
      onClose: function () {
        window.alert('window closed')
      }
    })
    handler.openIframe()
  }
}
