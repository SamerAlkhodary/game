import React, { useContext, useEffect, useState } from 'react';
import UserContext from '../../contexts/userContext';
import CartContext from '../../contexts/cartContext';

import { AppBar, Toolbar, CssBaseline, Typography, IconButton, Badge, CardMedia } from '@material-ui/core';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';

import { Link, useHistory } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import useStyles from './style';
import { AccountCircle } from '@material-ui/icons';
import HorizentalList from './HorizentalList';
import logo from '../../images/logo3.png'
const ShopPage = (props) => {
  const history = useHistory();

  const classes = useStyles();
  const { userState, load, logout } = useContext(UserContext);
  const { cartState, add } = useContext(CartContext);
  const StyledBadge = withStyles((theme) => ({
    badge: {
      right: -3,
      top: 13,
      border: `2px solid ${theme.palette.background.paper}`,
      padding: '0 4px',
    },
  }))(Badge);
  
  useEffect(() => {
    load();
  }, [load]);
  useEffect(() => {
    console.log("here", userState)
  }, [userState])
  const shoesData = [
    { id: "0", name: " Nike air", price: 300, unit: "SEK", img: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/04f7ca5f-2412-4511-993c-2e08e542eb33/air-max-90-shoes-6SdNzK.png", img1: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/3e63be82-043b-4eb1-988e-08769ab752f4/air-max-90-shoes-6SdNzK.png", type: "Shoes" },
    { id: "1", name: " Nike normal", price: 750, unit: "SEK", img: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/b82bf450-359e-4020-a0e8-5a7378e4bbbe/metcon-7-training-shoes-rNLHqX.png", img1: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/58e93ea3-50b6-4b0b-8703-ea0e496e21a3/metcon-7-training-shoes-rNLHqX.png", type: "Shoes" },
    { id: "2", name: " Nike black", price: 790, unit: "SEK", img: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/038a0be7-1fb7-441d-902f-98176a02ad81/react-infinity-run-flyknit-2-running-shoes-2GTTcm.png", img1: "https://static.nike.com/a/images/t_PDP_864_v1/f_auto,b_rgb:f5f5f5/19f25f1e-06fd-4bea-b07b-f1db05d472aa/react-infinity-run-flyknit-2-running-shoes-2GTTcm.png" },
    { id: "3", name: " Calvin Klein brown", price: 900, unit: "SEK", type: "Shoes", img: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00026_YAF_main?$main@2x$", img1: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00026_YAF_alternate3?$main@2x$", },
    { id: "4", name: " Calvin Klein red", price: 600, unit: "SEK", type: "Shoes", img: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00286_BEH_main?$main@2x$", img1: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00286_BEH_alternate3?$main@2x$" },
    { id: "5", name: " Calvin Klein blue", price: 1000, unit: "SEK", type: "Shoes", img: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00040_YAF_main?$main@2x$", img1: "https://calvinklein-eu.scene7.com/is/image/CalvinKleinEU/YM0YM00040_YAF_alternate4?$main@2x$" },
    { id: "6", name: " Ecco brown", price: 700, unit: "SEK", type: "Shoes", img: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dw7c08f1a5/productimages/470404-60207-outside.jpg?sw=1400&sh=1400&sm=fit&q=75", img1: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dwd9aba287/productimages/470404-60207-top.jpg?sw=1400&sh=1400&sm=fit&q=75" },
    { id: "7", name: " Ecco green", price: 800, unit: "SEK", type: "Shoes", img: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dw4259193e/productimages/512274-02559-main.jpg?sw=1400&sh=1400&sm=fit&q=75", img1: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dw06e19873/productimages/512274-02559-pair.jpg?sw=1400&sh=1400&sm=fit&q=75" },
    { id: "8", name: " Ecco three", price: 550, unit: "SEK", type: "Shoes", img: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dw243f9d88/productimages/450444-53779-outside.jpg?sw=1400&sh=1400&sm=fit&q=75", img1: "https://us.ecco.com/dw/image/v2/BCNL_PRD/on/demandware.static/-/Sites-ecco/default/dwc7f8c47d/productimages/450444-53779-top.jpg?sw=1400&sh=1400&sm=fit&q=75" }
  ];

  const perfumes = [
    { id: "1", name: " Carolina Herrera ch", price: 1000, unit: "SEK", img: "https://www.carolinaherrera.com/cdn-cgi/image/q=80,format=auto,fit=contain,width=1400/medias/sys_master/images/images/h00/h13/9131429265438/9131429265438.jpg", img1: "https://www.pricerunner.se/product/1200x630/3000211363/Carolina-Herrera-CH-for-Men-EdT-50ml.jpg" },
    { id: "2", name: " Black XS", price: 1100, unit: "SEK", img: "https://www.pacorabanne.com/cdn-cgi/image/fit=contain,width=1400,quality=90/medias/sys_master/images/images/hf3/h87/9242390396958/9242390396958.jpg", img1: "https://www.pacorabanne.com/cdn-cgi/image/fit=contain,width=1400,quality=90/medias/sys_master/images/images/h66/h22/9242391379998/9242391379998.jpg" },
    { id: "3", name: " One Million", price: 1150, unit: "SEK", img1: "https://www.pacorabanne.com/cdn-cgi/image/fit=contain,width=1400,quality=90/medias/sys_master/images/images/h78/hf4/9251973365790/9251973365790.jpg", img: "https://www.pacorabanne.com/cdn-cgi/image/fit=contain,width=1200,quality=90/medias/sys_master/images/images/hae/hd6/9251974283294/9251974283294.jpg" },
    { id: "4", name: " Chanel", price: 1200, unit: "SEK", img: "https://www.chanel.com/images//t_one/t_fragrance//q_auto:good,f_jpg,fl_lossy,dpr_1.2/w_1920/bleu-de-chanel-parfum-spray-3-4fl-oz--packshot-default-107180-8841593094174.jpg", img1: "https://m.media-amazon.com/images/I/41xO+-LGHAL._AC_.jpg" },
    { id: "5", name: " Fahrenheit", price: 900, unit: "SEK", img: "https://www.dior.com/beauty/version-5.1595261547035/resize-image/ep/715/773/90/0/horizon%252Fimages_additionnelles%252FY0096230_E03_GHC.jpg", img1: "https://cdn.fragrancenet.com/images/photos/900x900/288245.jpg" },
  ]

  return (
    <div className={classes.root}>
      <CssBaseline />

      <AppBar position="static" className={classes.appBar} elevation={1}>
        <Toolbar>

          <img src={logo} alt={"hi"} className={classes.title}/>
      
          <div className={classes.space}></div>

          <IconButton onClick={() => { history.push('/cart') }}>
            <StyledBadge badgeContent={cartState.items.length} color="secondary" className={classes.icon}>
              <ShoppingCartIcon style={{ color: 'black' }} />
            </StyledBadge>
          </IconButton>
          {
            userState.token ? <IconButton edge="end" aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={logout}
            >
              <AccountCircle style={{ color: 'black' }} />
            </IconButton> : <Link to="/login"><IconButton ><div><Typography variant="h6" className={classes.loginButton}>Login</Typography></div></IconButton></Link>
          }
        </Toolbar>
      </AppBar>
      <main className={classes.root}>
        <HorizentalList data={shoesData} title="Shoes" add={add} />
        <HorizentalList data={perfumes} title="Fragrances" add={add} />
        <HorizentalList data={shoesData} title="Snickers" add={add} />
        <HorizentalList data={perfumes} title="Perfumes" add={add} />
      </main>
    </div>

  );
}

export default (ShopPage);