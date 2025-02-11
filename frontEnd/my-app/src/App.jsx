import './App.css'
import {BrowserRouter, Routes, Route} from 'react-router-dom'
import Navbar from './Component/Navbar/Navbar'
import Shop from './Pages/Shop'
import ShopCatagory from './Pages/ShopCatagory'
import Product from './Pages/Product'
function App() {
  

  return (
    <div>
      <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Shop />} />
        <Route path="/mens" element={<ShopCatagory category="men" />} />
        <Route path="/womens" element={<ShopCatagory category="women"/>} />
        <Route path="/kids" element={<ShopCatagory category="kid" />} />
        <Route path='product' element={<Product/>}> 
          <Route path=':productID' element={<Product/>}/>
        </Route>
      </Routes>
      </BrowserRouter>
      
    </div>
  )
}

export default App
