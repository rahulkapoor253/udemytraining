
public class Maruti extends Cars{
	 //int resale_val;//60% of price
	
	Maruti(String name, int price)
	{
		this.model = name; 
		this.price = price;
		resale_val = this.price*60/100;
	}
		
	
	
}
