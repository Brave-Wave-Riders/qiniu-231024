export const validatePassword = () => {
  return (rule, value, callback) => {
    if (value.length < 6) {
      callback(new Error('pwd error'))
    } else {
      callback()
    }
  }
}
