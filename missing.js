(() => {
  let x = '';
  let data = {};
  document.querySelectorAll('a').forEach(c => {
    let year = c.href.split("/")[3];
    if (!(year in data)) {
      data[year] = {};
    }
    data[year][c.href.split("/").slice(3, 6).join("-")] = true;
  });

  const currentDate = new Date();

  for (const year in data) {
    const dates = Object.keys(data[year]);
    const yearInt = parseInt(year);
    if (isNaN(yearInt)) continue;
    const isLeapYear = (yearInt % 4 === 0 && yearInt % 100 !== 0) || yearInt % 400 === 0;
    const daysInYear = isLeapYear ? 366 : 365;
    const fullYear = new Date(yearInt, 0, 1).getFullYear();
    x += `Missing dates for ${fullYear}:\n`;
    const loopEnd = yearInt === currentDate.getFullYear() ? currentDate.getDate() : daysInYear;
    for (let i = 1; i <= loopEnd; i++) {
      const currentDate = new Date(yearInt, 0, i + 1);
      const dateString = currentDate.toISOString().slice(0, 10);
      if (!dates.includes(dateString)) {
        x += dateString + '\n';
      }
    }
  }
  console.log(x);
})();

