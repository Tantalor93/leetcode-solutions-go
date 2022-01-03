/*
There are n flights that are labeled from 1 to n.

You are given an array of flight bookings bookings, where bookings[i] = [firsti, lasti, seatsi] represents a booking for flights firsti through lasti (inclusive) with seatsi seats reserved for each flight in the range.

Return an array answer of length n, where answer[i] is the total number of seats reserved for flight i.
*/

func corpFlightBookings(bookings [][]int, n int) []int {
    var res []int = make([]int, n)
    for _,v := range bookings {
        for i := v[0]; i <= v[1]; i++ {
            res[i-1]+=v[2]
        }
    }
    return res
    
}
