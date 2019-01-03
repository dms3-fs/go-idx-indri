

//
// MetadataPairVector.i - new, see MetadataPair.i for tds code on metadata pair.
//

#ifdef SWIGJAVA
#endif

#ifdef SWIGCSHARP
#endif

#ifdef SWIGGO

/* -----------------------------------------------------------------------------
 * std_vector.i - metadata pair specialization
 * ----------------------------------------------------------------------------- */

%{
#include <vector>
#include <stdexcept>
%}

namespace std {

    // MetadataPair specialization
    template<> class vector<indri::parse::MetadataPair> {
      public:
        typedef size_t size_type;
        typedef indri::parse::MetadataPair value_type;
        typedef indri::parse::MetadataPair const_reference;
        vector();
        vector(size_type n);
        size_type size() const;
        size_type capacity() const;
        void reserve(size_type n);
        %rename(isEmpty) empty;
        bool empty() const;
        void clear();
        %rename(add) push_back;
        void push_back(const value_type& x);

        %extend {
            value_type get(int i) throw (std::out_of_range) {
                int size = int(self->size());
                if (i>=0 && i<size)
                    return (*self)[i];
                else
                    throw std::out_of_range("vector index out of range");
            }
            void set(int i, const value_type& val) throw (std::out_of_range) {
                int size = int(self->size());
                if (i>=0 && i<size)
                    (*self)[i] = val;
                else
                    throw std::out_of_range("vector index out of range");
            }
        }

    };

    %template(MetadataPairVector) vector<indri::parse::MetadataPair>;

}

#endif
