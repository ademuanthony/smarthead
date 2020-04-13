import { Controller } from 'stimulus'
import axios from 'axios'
import { hide, show } from '../utils'
import _ from 'lodash-es'

export default class extends Controller {
  list
  periods
  subjects
  classes

  static get targets () {
    return [
      'class', 'subject', 'period', 'cartItemDiv', 'listTbl', 'itemTemplate', 'cartTotal', 'savings'
    ]
  }

  connect () {
    this.list = []
    this.periods = []
    this.classes = []
    this.subjects = []
    const that = this
    Array.prototype.forEach.call(this.periodTarget.options, function (opt) {
      if (opt.value === '') return
      that.periods.push({
        id: opt.value,
        label: opt.innerText
      })
    })
    Array.prototype.forEach.call(this.subjectTarget.options, function (opt) {
      if (opt.value === '') return
      that.subjects.push({
        id: opt.value,
        label: opt.innerText
      })
    })
    Array.prototype.forEach.call(this.classTarget.options, function (opt) {
      if (opt.value === '') return
      that.classes.push({
        id: opt.value,
        label: opt.innerText
      })
    })
  }

  addToList () {
    const subjectID = this.subjectTarget.value
    const periodID = this.periodTarget.value
    const classID = this.classTarget.value
    if (subjectID === '' || periodID === '' || classID === '') {
      window.alert('Subject, period amd class are required')
      return
    }

    let id = subjectID + periodID + classID
    if (_.find(this.list, function (item) {
      return item.id === id
    })) {
      window.alert('Select subject exists for the specified period and class')
      return
    }

    this.list.push({
      id: id,
      subject_id: subjectID,
      period_id: periodID,
      class_id: classID,
      subject: this.subjectTarget.options[this.subjectTarget.selectedIndex].innerText,
      period: this.periodTarget.options[this.periodTarget.selectedIndex].innerText,
      className: this.classTarget.options[this.classTarget.selectedIndex].innerText
    })

    this.displayList()
  }

  removeFromList (evt) {
    let id = evt.currentTarget.getAttribute('data-id')
    _.remove(this.list, function (item) {
      return item.id === id
    })
    this.displayList()
  }

  displayList () {
    const _this = this
    this.listTblTarget.innerHTML = ''
    this.cartTotal = 0

    this.list.forEach((item, i) => {
      const exRow = document.importNode(_this.itemTemplateTarget.content, true)
      const fields = exRow.querySelectorAll('td')

      fields[0].innerText = i + 1
      fields[1].innerText = item.className
      fields[2].innerText = item.subject
      fields[3].innerHTML = item.period
      fields[4].innerHTML = `<button data-action="click->payment#removeFromList" data-id="${item.id}">Remove</button>`

      _this.listTblTarget.appendChild(exRow)
    })

    let count = this.list.length
    if (count >= 5) {
      let fives = (count - count % 5) / 5
      this.cartTotal = 15000 * fives
      count -= fives * 5
    }
    if (count >= 3) {
      let threes = (count - count % 3) / 3
      this.cartTotal += 12000 * threes
      count -= threes * 3
    }
    this.cartTotal += (count * 5000)

    this.cartTotalTarget.textContent = this.cartTotal
    let savings = (this.list.length * 5000) - this.cartTotal
    let savingsPercentage = 100 * savings / this.cartTotal

    this.savingsTarget.textContent = `${savings} (${savingsPercentage.toFixed(0)}%)`
    if (this.list.length > 0) {
      show(this.cartItemDivTarget)
    } else {
      hide(this.cartItemDivTarget)
    }
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
      'class_id': this.list[0].class_id,
      'period_id': this.list[0].period_id
    }
    const that = this
    const resp = await axios.post('/payments/initiate', req)
    const result = resp.data
    if (result.error) {
      window.alert(result.error)
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
            'class_id': that.list[i].class_id,
            'period_id': that.list[i].period_id
          })
        }
        const resp = await axios.post(`/payments/${response.reference}/update-status`, { 'items': items })
        const result = resp.data
        if (result.error) {
          window.alert(result.error)
        }
        window.alert('Success. Subscription successful, check the subscription table for your lesson\' start date')
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
