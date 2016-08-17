require 'httparty'

class List
  def self.get_all()
    url = "http://localhost:3002"
    response = HTTParty.get(url+"/list/all")
    response.parsed_response
  end

  def self.get_list(id)
    url = "http://localhost:3002"
    response = HTTParty.get(url+"/lists/#{id}/show")
    response.parsed_response
  end

end
