package calendar_test

import (
	. "github.com/tn5993/projecteuler/utils/calendar"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calendar", func() {
	Context("Test a specific year is a leap year", func() {
		It("should return true or false if a year is a leap year", func() {
			Ω(IsLeapYear(2012)).To(BeTrue())
			Ω(IsLeapYear(2016)).To(BeTrue())
			Ω(IsLeapYear(2100)).To(BeFalse())
			Ω(IsLeapYear(2200)).To(BeFalse())
			Ω(IsLeapYear(2000)).To(BeTrue())
			Ω(IsLeapYear(2400)).To(BeTrue())
		})
	})

	Context("Test get doomday of year", func() {
		It("should return the correct doomday of year", func() {
			Ω(GetDoomDayOfYear(1981)).To(Equal(6))
			Ω(GetDoomDayOfYear(1980)).To(Equal(5))
			Ω(GetDoomDayOfYear(1985)).To(Equal(4))
			Ω(GetDoomDayOfYear(1979)).To(Equal(3))
			Ω(GetDoomDayOfYear(1978)).To(Equal(2))
			Ω(GetDoomDayOfYear(1977)).To(Equal(1))
			Ω(GetDoomDayOfYear(1976)).To(Equal(0))
		})
	})

	Context("Get day of week", func() {
		It("should return the correct day of week", func() {
			Ω(GetDayOfWeek(2, 27, 1979)).To(Equal(2))
			Ω(GetDayOfWeek(2, 11, 1978)).To(Equal(6))
			Ω(GetDayOfWeek(2, 10, 1977)).To(Equal(4))
			Ω(GetDayOfWeek(10, 10, 1960)).To(Equal(1))
		})
	})
})
