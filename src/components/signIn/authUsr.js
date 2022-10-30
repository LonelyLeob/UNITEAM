// import axios from "axios";
//
//  async function authUsr(url, data){
//
//
//     let json = JSON.stringify({"name": name, "password": password})
//     console.log(json)
//     // axios.post("http://localhost:7000/auth/token", json, {withCredentials: true}).then(data => console.log(data))
//
//
//
//      let response = await fetch(url,{
//          method:'POST',
//          body: JSON.stringify(data),
//      })
//      if (!response.ok){
//          throw new Error('Ошибка' )
//      }
//    return await response.json()
// }
//
// export default authUsr()