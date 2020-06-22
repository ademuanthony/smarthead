import { Controller } from 'stimulus'
import axios from 'axios'
import { hide, show } from '../utils'
import _ from 'lodash-es'

export default class extends Controller {
  list
  subjects
  classes

  static get targets () {
    return [
      'regNo', 'startDate', 'endDate',
      'subject', 'cartItemDiv', 'listTbl', 'itemTemplate', 'cartTotal', 'savings'
    ]
  }

  connect () {
    this.list = []
    this.subjects = []
    const that = this
    Array.prototype.forEach.call(this.subjectTarget.options, function (opt) {
      if (opt.value === '') return
      that.subjects.push({
        id: opt.value,
        label: opt.innerText
      })
    })
  }

  addToList () {
    const subjectID = this.subjectTarget.value
    if (subjectID === '') {
      window.alert('Subject is required')
      return
    }

    let id = subjectID
    if (_.find(this.list, function (item) {
      return item.subject_id === id
    })) {
      window.alert('Select subject exists')
      return
    }

    this.list.push({
      subject_id: subjectID,
      subject: this.subjectTarget.options[this.subjectTarget.selectedIndex].innerText
    })

    this.displayList()

    this.subjectTarget.value = ''
  }

  removeFromList (evt) {
    let id = evt.currentTarget.getAttribute('data-id')
    _.remove(this.list, function (item) {
      return item.subject_id === id
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
      fields[1].innerText = item.subject
      fields[2].innerHTML = `<button data-action="click->subscription#removeFromList" data-id="${item.id}">Remove</button>`

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

  async save (e) {
    e.preventDefault()
    if (this.list.length === 0) {
      window.alert('Please add a subject')
      return
    }

    if (this.loading) return
    this.loading = true

    const req = {
      'student_reg_no': this.regNoTarget.value,
      'start_date': this.startDateTarget.value,
      'end_date': this.endDateTarget.value,
      'items': this.list.map(item => {
        return { subject_id: item.subject_id }
      })
    }
    const that = this
    axios.post('/admin/subscriptions/create', req).then(() => {
      window.location.href = '/admin/subscriptions'
    }).catch(err => {
      let error = err.response.data.details
      that.loading = false
      console.log(err)
      window.alert(error)
    })
  }
}
