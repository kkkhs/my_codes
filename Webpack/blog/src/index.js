import { getBlogPosts } from "./data";
import "./style.css"
import img1 from './asset/imgs/2.jpeg'   //导入图片

const blogs = getBlogPosts()
const ul = document.createElement('ul')
blogs.forEach((blog) => {
  const li = document.createElement('li')
  li.innerHTML = blog
  ul.appendChild(li)
})
document.body.appendChild(ul)

const image = document.createElement('img')
image.src = img1
document.body.prepend(image)

document.body.prepend(document.createElement('h1').innerHTML='博客列表')
