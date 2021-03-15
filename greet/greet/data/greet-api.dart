// --greet--

class Request{
	
	/// 
	final {string} name;
	
	Request({ 
		this.name,
	});
	factory Request.fromJson(Map<String,dynamic> m) {
		return Request(
			name: 