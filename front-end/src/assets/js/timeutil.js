// date1 and date2 are javascript Date objects, return day2 - day1
export function dateDiffInDays(date1, date2) {
  const _MS_PER_DAY = 1000 * 60 * 60 * 24;
  // Discard the time and time-zone information.
  const utc1 = Date.UTC(date1.getFullYear(), date1.getMonth(), date1.getDate());
  const utc2 = Date.UTC(date2.getFullYear(), date2.getMonth(), date2.getDate());

  return Math.floor((utc2 - utc1) / _MS_PER_DAY);
}

export function getWeekday(date) {
  var days = ['Chủ Nhật','Thứ hai','Thứ ba','Thứ tư','Thứ năm','Thứ sáu','Thứ bảy'];

  // Discard the time and time-zone information.
  return days[ date.getDay() ];
}

export function dateTimeToString(date, format) {
  const offset = date.getTimezoneOffset();
  date = new Date(date.getTime() - (offset*60*1000));
  const dateString = date.toISOString().split('T')[0];
  const timeString = date.toISOString().split('T')[1];
  const day = dateString.split('-')[2];
  const month = dateString.split('-')[1];
  const year = dateString.split('-')[0];
  const hour = timeString.split(':')[0];
  const minute = timeString.split(':')[1];
  return format.toLowerCase().replace("dd", day).replace("mm", month).replace("yyyy", year).replace("hh", hour).replace("mi", minute);
}