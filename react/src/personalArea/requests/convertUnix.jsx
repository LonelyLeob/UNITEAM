function ConvertUnix(unix){
    var t = new Date(unix * 1000);
    let day = t.getDate()
    let month = t.getMonth() + 1
    let year = t.getFullYear()

    let hours = t.getHours()
    let minutes = t.getMinutes()

    return `${day}.${month}.${year} ${hours}:${minutes}`
}

export default ConvertUnix