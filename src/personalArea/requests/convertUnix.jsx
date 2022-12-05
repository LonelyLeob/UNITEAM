
function ConvertUnix(unix){
    let date = new Intl.DateTimeFormat('ru-RU', { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', second: '2-digit' }).format(unix)
    return date
}

export default ConvertUnix