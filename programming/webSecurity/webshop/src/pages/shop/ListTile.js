import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import React, { useContext, useEffect, useState } from 'react';
import { GridList, Typography,GridListTile,GridListTileBar,IconButton} from '@material-ui/core';
import useStyles from './style';

const ListTile=({tile, add})=>{
    const classes = useStyles();
    const [hoverState, setHover] = useState(false)  
    return (

     <GridListTile key={tile.id} className={classes.tile} onMouseOver={()=>{setHover(true)}} onMouseOut={() => setHover(false)} >
         <img src={hoverState?tile.img1:tile.img} alt={tile.title} />
     <GridListTileBar
       title={tile.price +" "+ tile.unit}
       classes={{
         root: classes.titleBar,
         title: classes.tileTitle,
       }}
       actionIcon={
         <IconButton aria-label={`star ${tile.title}`} onClick={()=>add(tile)
         }>
           <ShoppingCartIcon className={classes.tileTitle} />
         </IconButton>
       }
     />
   </GridListTile>

    ) 
}
export default ListTile;