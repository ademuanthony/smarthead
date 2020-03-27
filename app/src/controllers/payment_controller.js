import { Controller } from 'stimulus'
import axios from 'axios'

export default class extends Controller {
  static get targets () {
    return [
      'navbar'
    ]
  }

  async initiate () {
    const resp = await axios.get('/payments/initiate')
    const result = resp.data
    if (result.error) {
      window.alert(result.error)
      return
    }
    // eslint-disable-next-line no-undef
    var handler = PaystackPop.setup({
      key: 'pk_test_6e0424ff8e08138c2fcce3f4f2b05052c4b3f77c',
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
        window.alert('success. transaction ref is ' + response.reference)
        const resp = await axios.get(`/payments/${response.reference}/update-status`)
        const result = resp.data
        if (result.error) {
          window.alert(result.error)
        }
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
