import { GridList, Typography,GridListTile,GridListTileBar,IconButton} from '@material-ui/core';
import withWidth, { isWidthUp } from '@material-ui/core/withWidth';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import React, { useContext, useEffect, useState } from 'react';
import ListTile from './ListTile';

import useStyles from './style';

const HorizentalList= ({title,width,data, add })=>{
    
    const classes = useStyles();
    const listTiles = data.map((tile) =>
        <ListTile tile={tile} add={add}/>
      );

    function getCols(screenWidth) {
        if (isWidthUp('lg', screenWidth)) {
          return 4;
        }
    
        if (isWidthUp('md', screenWidth)) {
          return 3;
        }
    
        return 2;
      }
   return(
    <div className={classes.horizentalList}>
    <Typography align="center" gutterBottom variant="h4">{title}</Typography>
    <div className={classes.list}>
      <GridList className={classes.gridList} cols={getCols(width)} spacing={10}>
        {
          listTiles
        }
      </GridList>
    </div>
    </div>
   );
}
export default withWidth() (HorizentalList);